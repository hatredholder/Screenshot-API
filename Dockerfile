# building the screenshot app
FROM golang:alpine as build

WORKDIR /src

COPY . .

RUN go build -o screenshot .

# installing chromedp (better chromium driver)
FROM chromedp/headless-shell:stable

# launching the app
WORKDIR /app
COPY --from=build /src/screenshot ./

EXPOSE 8080

ENTRYPOINT ["/app/screenshot"]
