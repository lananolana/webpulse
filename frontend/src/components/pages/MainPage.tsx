import { useEffect, useState } from 'react';
import logo from '../../assets/images/logo.svg';
import styles from './MainPage.module.scss';
import { LinkButton } from '../UI/LinkButton/LinkButton';
import { Form } from '../UI/Form/Form';
import { useDispatch } from 'react-redux';
import { AppDispatch, useAppSelector } from '../../utils/hooks';
import { getInfo } from '../../services/mainSlice';

function MainPage() {
  const dispatch = useDispatch<AppDispatch>();
  const [firstRender, setFirstRender] = useState<boolean>(true);
  const errorMessage = useAppSelector((state) => state.data.errorMessage);

  function onClick(e: React.MouseEvent, text:string) {
    e.preventDefault();
    if (firstRender) {
      setFirstRender(false)
    }
    dispatch(getInfo(text));
  }

  return (
    <main className={styles.main}>
      {firstRender ? 
      <>
        <div className={styles.logo}>
          <img src={logo} alt="Логотип"/>
          <h1 className={styles.logo__title}>
            WEBPULSE
          </h1>
        </div>
        <Form onClick={onClick}/>
      </> : 
      <div className={styles.topRow}>
      <img src={logo} alt="Логотип"/>
      <Form onClick={onClick}/>
      </div>}
      <div className={styles.more_often}>
        {!errorMessage ? 
        <>
        <h2 className={styles.more_often__title}>Чаще всего ищут</h2>
          <LinkButton 
            text={'google.com'}
            onClick={onClick}
          />
          <LinkButton 
            text={'telegram.ru'}
            onClick={onClick}
          />
          <LinkButton 
            text={'yandex.com'}
            onClick={onClick}
          />
        </> : (errorMessage !== null) &&
        <p className={styles.error__text}>
          Упс! Что-то пошло не так. Попробуйте ещё раз.
        </p>}
      </div>    
    </main>
  )
}

export default MainPage;