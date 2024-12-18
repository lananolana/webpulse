import { useState } from 'react';
import logo from '../../assets/images/logo.svg';
import styles from './MainPage.module.scss';
import { LinkButton } from '../UI/LinkButton/LinkButton';
import { Form } from '../UI/Form/Form';
import { useDispatch } from 'react-redux';
import { AppDispatch, useAppSelector } from '../../utils/hooks';
import { getInfo } from '../../services/mainSlice';
import ResBlock from '../UI/ResBlock/ResBlock';

function MainPage() {
  const dispatch = useDispatch<AppDispatch>();
  const [firstRender, setFirstRender] = useState<boolean>(true);
  const errorMessage = useAppSelector((state) => state.data.errorMessage);
  const [form, setForm] = useState('');
  const data = useAppSelector((state) => state.data.data);

  function onClick(e: React.MouseEvent, text:string) {
    e.preventDefault();
    if (firstRender) {
      setFirstRender(false)
    }
    dispatch(getInfo(text));
    setForm(text)
  }

  return (
    <main className={styles.main}>
      {firstRender ? 
      <>
        <div className={styles.logo}>
          <img src={logo} alt="Логотип"/>
        </div>
        <Form 
          onClick={onClick}
          form={form}
          setForm={setForm}/>
      </> : 
      <div className={styles.topRow}>
      <img src={logo} alt="Логотип"/>
      <Form 
        onClick={onClick}
        form={form}
        setForm={setForm}/>
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
            text={'yandex.ru'}
            onClick={onClick}
          />
        </> : (errorMessage !== null) &&
        <p className={styles.error__text}>
          Упс! Что-то пошло не так. Попробуйте ещё раз.
        </p>}
      </div>
      {!firstRender && !errorMessage && data &&
        <ResBlock />
      }
    </main>
  )
}

export default MainPage;