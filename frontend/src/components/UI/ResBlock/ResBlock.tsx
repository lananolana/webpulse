import { useEffect, useState } from 'react';
import size from '../../../assets/images/size.png';
import speed from '../../../assets/images/speed.png';
import status from '../../../assets/images/status.png';
import dns from '../../../assets/images/dns.png';
import security from '../../../assets/images/security.png';
import styles from './ResBlock.module.scss';
import { useAppSelector } from '../../../utils/hooks';
import { ResCard } from './ResCard';

const ResBlock = () => {
  const data = useAppSelector((state) => state.data.data);

  return (
    <>
    {data.status === "Success" ?
    <div className={styles.results}>
      <div className={styles.cardsContainer}>
        <ResCard 
          count={data?.availability?.http_status_code?.toString() ?? 'Нет данных'}
          text={data?.availability?.http_status_code === 200 ? 'Сайт доступен' : 'Сайт недоступен'}
          src={status}
        />
        <ResCard 
          count={data?.security?.ssl?.valid ? 'Безопасное соединение' : 'Небезопасное соединение'}
          text={data?.security?.cors ? 'CORS включён' : 'CORS не включён'}
          moreText={data?.security?.ssl?.expires_at ? 
            `Сертификат истекает ${new Date(data?.security?.ssl?.expires_at*1000).toLocaleString().slice(0, 10)}` : 
            'Нет данных'}
          src={security}
          type={'text'}
        />
        <ResCard 
          count={data?.performance?.response_time_ms?.toString() ?? 'Нет данных'}
          span={'мс'}
          text={'Время отклика ⓘ'}
          title={'Идеальное значение для полной загрузки HTML — менее 300 мс'}
          src={speed}
        />
        <ResCard
          count={data?.performance?.response_size_kb?.toString() ?? 'Нет данных'}
          span={'кб'}
          text={'Размер ответа ⓘ'}
          title={'Идеальное значение для быстрого времени загрузки — до 100 Kб'}
          src={size}
        />
        <ResCard
          count={data?.server_info?.dns_response_time_ms?.toString() ?? 'Нет данных'}
          span={'мс'}
          text={'DNS-проверка ⓘ'}
          title={'Идеальное значение для времени отклика DNS — менее 50 мс'}
          src={dns}
        />
      </div>
      <div className={styles.smallCardsContainer}>
        <ResCard
          count={data?.server_info?.ip_address ?? 'Нет данных'}
          text={'IP-адрес'}
          type={'small'}
        />
        <ResCard
          count={data?.server_info?.web_server ?? 'Нет данных'}
          text={'Тип веб-сервера'}
          type={'small'}
        />
        <ResCard
          count={data?.performance?.optimization ?? 'Нет данных'}
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