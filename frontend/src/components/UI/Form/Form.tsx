import { FC, useEffect, useRef, useState } from 'react';
import styles from './Form.module.scss';
import search from '../../../assets/images/search.svg';

type Props = {
  onClick(e: React.MouseEvent, text: string): void;
};

const Form: FC<Props> = ({
  onClick
}) => {
  const inputRef = useRef<HTMLInputElement | null>(null);
  const buttonRef = useRef<HTMLButtonElement | null>(null);
  const [form, setForm] = useState('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm(e.target.value);
  };

  return (
    <form className={styles.form}>
      <img src={search} alt="Иконка лупы" className={styles.form_img}/>
      <div className={styles.inputContainer}>
      <input
        name="search" 
        placeholder="Введите адрес сайта" 
        type="search"
        value={form}
        id='cardsHolderRepeatingInput'
        ref={inputRef}
        onChange={onChange}
        className={styles.form_search}
        maxLength={200}
      />
      {form !== '' && !/^((http|https):\/\/)?[a-zа-я0-9]+([\-\.]{1}[a-zа-я0-9]+)*\.[a-zа-я]{2,5}(:[0-9]{1,5})?(\/.*)?$/i.test(form) ? 
      <p className={styles.error_text}>Введён несуществующий URL</p> : form.length > 200 ? 
      <p className={styles.error_text}>URL не может превышать 200 символов</p> : <div style={{height: '68px'}}></div>}
      </div>
      <button 
        className={styles.form_button} 
        type="submit"
        ref={buttonRef}
        disabled={!/^((http|https):\/\/)?[a-zа-я0-9]+([\-\.]{1}[a-zа-я0-9]+)*\.[a-zа-я]{2,5}(:[0-9]{1,5})?(\/.*)?$/i.test(form) && form.length < 200}
        onClick={(e) => {if (form !== '' && 
          !inputRef.current?.validity.typeMismatch) onClick(e, form)}}>
        Проверить
      </button>
    </form>
  );
};

export { Form };