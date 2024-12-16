import { FC } from 'react';
import styles from './ResBlock.module.scss';
import classNames from 'classnames';

type Props = {
  count?: string;
  text?: string;
  moreText?: string;
  title?: string;
  src?: string;
  type?: string;
  span?: string
};

const ResCard: FC<Props> = ({
  count,
  text,
  moreText,
  title,
  src,
  type = 'big',
  span
}) => {

  return (
    <>
    {(type === 'big') ?
    <div className={styles.background}>
    <div className={classNames(styles.border, styles.card)} >
      <p className={styles.digit}>
        {count ? count : '???'}<span className={styles.smallDigit}>{span && span}</span>
      </p>
      <img src={src} alt="Иконка" className={styles.icon}/>
      <p title={title} className={styles.signification}>{text ? text : 'Нет информации'}</p>
    </div>
    </div> : (type === 'text') ?
    <div className={styles.background}>
    <div className={classNames(styles.border, styles.textCard)} >
      <p className={styles.smallDigit}>{count ? count : '???'}</p>
      <img src={src} alt="Иконка" className={styles.icon}/>
      <div>
      <p title={title} className={styles.signification}>
        {text && text}
      </p>
      <p title={title} className={styles.signification}>
        {moreText ? moreText : 'Нет информации'}
      </p>
      </div>
    </div></div>
    : (type === 'small') &&
    <div className={styles.background}>
    <div className={classNames(styles.border, styles.smallCard)} >
      <p className={styles.smallDigit}>{count ? count : '???'}</p>
      <p title={title} className={styles.signification}>{text ? text : 'Нет информации'}</p>
    </div> 
    </div>
    }
    </>
  );
};

export { ResCard };