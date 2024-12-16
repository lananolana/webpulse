import { FC } from 'react';
import s from './LinkButton.module.scss';

type Props = {
  text: string;
  onClick(e: React.MouseEvent, text: string): void;
};

const LinkButton: FC<Props> = ({
  text,
  onClick
}) => {

  return (
    <button
      className={s.button}
      onClick={(e: React.MouseEvent) => onClick(e, text)}
      type="button"
    >
      {text}
    </button>
  )
}

export { LinkButton };