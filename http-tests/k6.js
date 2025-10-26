// https://grafana.com/docs/k6/latest/results-output/web-dashboard/

// command: K6_WEB_DASHBOARD=true K6_WEB_DASHBOARD_OPEN=true K6_WEB_DASHBOARD_PERIOD=2s K6_WEB_DASHBOARD_EXPORT=html-report.html k6 run http-tests/k6.js
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    discardResponseBodies: true,
    scenarios: {
      contacts: {
        executor: 'ramping-vus',
        startVUs: 0,
        stages: [
          { duration: '30s', target: 500 },
        ],
        gracefulRampDown: '0s',
      },
    },
  };

export default function () {
    http.get('http://localhost:80/v1/user/99');
    sleep(0.5);

}