import { FC, useRef, useState } from 'react';
import styles from './Form.module.scss';
import search from '../../../assets/images/search.svg';

type Props = {
  text?: string;
  link?: string;
  onClick(e: React.MouseEvent, text: string): void;
};

const Form: FC<Props> = ({
  onClick
}) => {

  const inputRef = useRef<HTMLInputElement | null>(null);
  const [form, setForm] = useState('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm(e.target.value);
  };

  return (
    <form className={styles.form} action="" method="get">
      <img src={search} alt="Иконка лупы" className={styles.form_img}/>
      <input
        name="search" 
        placeholder="Введите адрес сайта" 
        type="search"
        value={form}
        id='cardsHolderRepeatingInput'
        ref={inputRef}
        onChange={onChange}
        className={styles.form_search}
      />
      <button 
        className={styles.form_button} 
        type="submit"
        onClick={(e) => {onClick(e, form)}}>
        Проверить
      </button>
    </form>
  );
};

export { Form };