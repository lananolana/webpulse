import { FC } from 'react';
import { Routes, Route } from 'react-router-dom';
import MainPage from './components/pages/MainPage';
import styles from './App.module.scss';

const App: FC = () => {
  return (
    <>
      <div className={styles.page}>
        <Routes>
          <Route path='*' element={<MainPage />}></Route>
        </Routes>
      </div>
    </>
  );
};

export default App;
