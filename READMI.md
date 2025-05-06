# Web Crawler
This application is an asynchronous crawler.
It parse web-pages, indentifies templates of popular web-builders (including tilda, bitrix, wordpress, html5)
    and returns structured blocks in JSON format.

# Features
- supports popular templates: tilda, bitrix, html5, wordpress;
- parse html elements: header, footer, image;
- has a conveniently expandable parser architecture;
- integration and unit tests;
- logging and secure data processing;

# Architecture
- internal/
    - crawler # implements a queue, worker pool, page crawling
    - dispatcher # centralized dispatcher for calling the required parser
    - model # block model and parser interfaces
    - parser # parsers based on templates
- pkg/
    - util/logger # wrap over Sugared logger

# Build and launch

### Docker
    '''bash
    docker build -t web-crawler .
    docker run -p 8080:8080 web-crawler

### Local build
    go build -o crawler ./cmd/server
    ./crawler

# Run crawler
    - Post request to http://host.com:port/crawl
    - Request structure:
        {"urls": [
            "https://example.com",
            "https://another.com"
            ]
        }

# Testing
    # Unit testing
        go test ./internal/... -v
    # Integration testing
        go test -tags=integration ./tests/...
    # Load testing
        docker run -i grafana/k6 run /scripts/loadtest.js < script.js
    
        