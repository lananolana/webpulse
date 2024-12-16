import { useEffect, useState } from 'react';
import size from '../../../assets/images/size.png';
import speed from '../../../assets/images/speed.png';
import status from '../../../assets/images/status.png';
import dns from '../../../assets/images/dns.png';
import security from '../../../assets/images/security.png';
import styles from './ResBlock.module.scss';
import { useDispatch } from 'react-redux';
import { AppDispatch, useAppSelector } from '../../../utils/hooks';
import { ResCard } from './ResCard';

function ResBlock() {
  const dispatch = useDispatch<AppDispatch>();
  const data = useAppSelector((state) => state.data.data);

  //Пока тут мок, но однажды будут данные
  const mock = {
    status: "Success",
    availability: {
      is_available: true,
      http_status_code: 200
    },
    performance: {
      response_time_ms: 614,
      transfer_speed_kbps: 184,
      response_size_kb: 13,
      optimization: "gzip"
    },
    security: {
      ssl: {
        valid: true,
        expires_at: 1749317007,
        issuer: "CN=GlobalSign RSA OV SSL CA 2018,O=GlobalSign nv-sa,C=BE"
      },
      cors: {
        enabled: true,
        allow_origin: "yastatic.net"
      }
    },
    server_info: {
      ip_address: "77.88.55.88",
      web_server: "gws",
      dns_response_time_ms: 83,
      dns_records: {
        A: [
          "77.88.55.88",
          "77.88.44.55",
          "5.255.255.77",
          "2a02:6b8:a::a"
        ],
        CNAME: "yandex.ru.",
        MX: [
          "mx.yandex.ru."
        ]
      }
    }
  }

  return (
    <>
    {mock.status === "Success" ?
    <div className={styles.results}>
      <div className={styles.cardsContainer}>
        <ResCard 
          count={mock?.availability?.http_status_code?.toString() ?? 'Нет данных'}
          text={mock?.availability?.http_status_code === 200 ? 'Сайт доступен' : 'Сайт недоступен'}
          src={status}
        />
        <ResCard 
          count={mock?.security?.ssl?.valid ? 'Безопасное соединение' : 'Небезопасное соединение'}
          text={mock?.security?.cors ? 'CORS включён' : 'CORS не включён'}
          moreText={`Сертификат истекает ${mock?.security?.ssl?.expires_at}`}
          src={security}
          type={'text'}
        />
        <ResCard 
          count={mock?.performance?.response_time_ms?.toString() ?? 'Нет данных'}
          span={'мс'}
          text={'Время отклика ⓘ'}
          title={'Идеальное значение для полной загрузки HTML — менее 300 мс'}
          src={speed}
        />
        <ResCard
          count={mock?.performance?.response_size_kb?.toString() ?? 'Нет данных'}
          span={'кб'}
          text={'Размер ответа ⓘ'}
          title={'Идеальное значение для быстрого времени загрузки — до 100 Kб'}
          src={size}
        />
        <ResCard
          count={mock?.server_info?.dns_response_time_ms?.toString() ?? 'Нет данных'}
          span={'мс'}
          text={'DNS-проверка ⓘ'}
          title={'Идеальное значение для времени отклика DNS — менее 50 мс'}
          src={dns}
        />
      </div>
      <div className={styles.smallCardsContainer}>
        <ResCard
          count={mock?.server_info?.ip_address ?? 'Нет данных'}
          text={'IP-адрес'}
          type={'small'}
        />
        <ResCard
          count={mock?.server_info?.web_server ?? 'Нет данных'}
          text={'Тип веб-сервера'}
          type={'small'}
        />
        <ResCard
          count={mock?.performance?.optimization ?? 'Нет данных'}
          text={'Оптимизация'}
          type={'small'}
        />
      </div>
    </div> :
    <p className={styles.error__text}>
      Упс! Похоже, с адресом что-то не так.
    </p>
    }
    </>
  )
}

export default ResBlock;