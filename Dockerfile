FROM syui/aios

RUN pacman -Syu --noconfirm go
WORKDIR /app
RUN git clone https://git.syui.ai/ai/api tmp

WORKDIR /app/tmp
RUN go build
RUN mv api /app/api

#ADD ./app /app

WORKDIR /app
ENTRYPOINT ["/app/api"]
