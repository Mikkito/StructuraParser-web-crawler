import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 50, // virtual user`s
    duration: '30s', // test duration
};

export default function() {
    const payload = JSON.stringify({
        urls: [
            "https://test.com",
            "https://test2.com"
        ]
    });

    const headers = {
        'Content-Type': 'application/json',
    };

    const res = http.post('http://localhost:8080/crawl', payload, { headers });

    check(res, {
        'status is 200': (r) => r.status === 200,
        'body is not empty': (r) => r.body.length > 0,
    });

    sleep(1); // pause between requests to the service
}