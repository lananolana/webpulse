import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 1000, // количество виртуальных пользователей
    duration: '10s', // длительность теста
};

const domains = [
    'google.com',
    'bing.com',
    'duckduckgo.com',
    'microsoft.com',
    'github.com',
    'reddit.com',
    'stackoverflow.com',
    'apple.com',
    'samsung.com',
    'drive.google.com',
    'tiktok.com',
    'uber.com',
    'dropbox.com',
    'ford.com',
    'bmw.com',
    'toyota.com',
    'ikea.com',
    'slack.com',
    'atlassian.com',
    'vimeo.com',
    'twitch.tv',
    'booking.com',
    'qatarairways.com',
    'lufthansa.com',
    'britishairways.com',
    'zoom.com',
    'outlook.com',
    'sears.com',
    'zara.com',
    'sephora.com',
    'loreal.com',
    'gillette.com',
    'dove.com',
    'burgerking.com',
    'k6.io',
    'docker.com',
    'figma.com',
    'ipinfo.io',
    'draw.io'
];

export default function () {
    const domain = domains[Math.floor(Math.random() * domains.length)];
    const url = `http://localhost:8080/api/status?domain=${domain}`;
    const res = http.get(url);

    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time is less than 5s': (r) => r.timings.duration < 5000,
    });

    sleep(0.1); // пауза между запросами: 1 пользователь генерирует 10 запросов в секунду
}
