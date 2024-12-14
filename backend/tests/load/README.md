# Нагрузочное тестирование

Если mock_response=false, лучше выставить какое-то адекватное количество VUs, чтобы не упереться на 429.

В случае, если mock_response=true тестируется только работа сервиса без обращения к таргета, соответственно, мы получаем
достаточно гигантский RPS, количество которого упирается в вычислительные мощности и лимит количества сетевых соединений.

## Результаты нагрузочного тестирования в реальных условиях (mock_response=false)
Конфигурация:
- mock_response=false
- VUs=80
- duration=10s
- sleep=0.1s (Время бездействия VU между запросами)

#### Summary
146.409933 RPS

```
✓ status is 200
     ✗ response time is less than 5s
      ↳  99% — ✓ 1696 / ✗ 4

     checks.........................: 99.88% 3396 out of 3400
     data_received..................: 1.3 MB 110 kB/s
     data_sent......................: 185 kB 16 kB/s
     http_req_blocked...............: avg=152.88µs min=3.86µs   med=8.02µs   max=8.49ms   p(90)=12.63µs  p(95)=95.71µs 
     http_req_connecting............: avg=66.67µs  min=0s       med=0s       max=6.27ms   p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=387.95ms min=34.82ms  med=200.19ms max=5.84s    p(90)=763.27ms p(95)=1.26s   
       { expected_response:true }...: avg=387.95ms min=34.82ms  med=200.19ms max=5.84s    p(90)=763.27ms p(95)=1.26s   
     http_req_failed................: 0.00%  0 out of 1700
     http_req_receiving.............: avg=91.78µs  min=24.14µs  med=83.4µs   max=1.08ms   p(90)=129.99µs p(95)=154.85µs
     http_req_sending...............: avg=43.13µs  min=8.77µs   med=26.62µs  max=987.32µs p(90)=43.14µs  p(95)=74.2µs  
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=387.82ms min=34.71ms  med=200.07ms max=5.84s    p(90)=763.03ms p(95)=1.26s   
     http_reqs......................: 1700   146.409933/s
     iteration_duration.............: avg=488.9ms  min=135.15ms med=300.92ms max=5.94s    p(90)=867.44ms p(95)=1.36s   
     iterations.....................: 1700   146.409933/s
     vus............................: 9      min=9            max=80
     vus_max........................: 80     min=80           max=80
```

## Результаты нагрузочного тестирования (mock_response=true)
Конфигурация:
- mock_response=true
- VUs=1000
- duration=10s
- sleep=0.1s (Время бездействия VU между запросами)

#### Summary
9397.102824 RPS

```
✓ status is 200
     ✓ response time is less than 5s

     checks.........................: 100.00% 189974 out of 189974
     data_received..................: 58 MB   5.8 MB/s
     data_sent......................: 10 MB   1.0 MB/s
     http_req_blocked...............: avg=475.29µs min=2.09µs   med=4.16µs   max=176.84ms p(90)=7µs      p(95)=8.74µs  
     http_req_connecting............: avg=455.46µs min=0s       med=0s       max=107.91ms p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=3.02ms   min=100.04µs med=908.34µs max=158.76ms p(90)=5.34ms   p(95)=10.93ms 
       { expected_response:true }...: avg=3.02ms   min=100.04µs med=908.34µs max=158.76ms p(90)=5.34ms   p(95)=10.93ms 
     http_req_failed................: 0.00%   0 out of 94987
     http_req_receiving.............: avg=332.11µs min=13.33µs  med=26.45µs  max=81.54ms  p(90)=140.9µs  p(95)=379.06µs
     http_req_sending...............: avg=337.51µs min=5.6µs    med=11.7µs   max=81.61ms  p(90)=229.64µs p(95)=837.25µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=2.35ms   min=75.82µs  med=780.68µs max=78.36ms  p(90)=4.33ms   p(95)=8.42ms  
     http_reqs......................: 94987   9397.102824/s
     iteration_duration.............: avg=105.01ms min=100.18ms med=101.57ms max=291.73ms p(90)=110.1ms  p(95)=119.35ms
     iterations.....................: 94987   9397.102824/s
     vus............................: 1000    min=1000             max=1000
     vus_max........................: 1000    min=1000             max=1000
```