import styles from './NavbarStyle.module.css'
import { FC } from 'react';
import { Link } from 'react-router-dom'
import { useAuth } from '../../context';
import Icons, { HomeIcon, ForwardArrowIcon } from '../../components/Icons';

const Navbar: FC = () => {
  const { login, logout } = useAuth()

  return (
    <div id={styles['navbar']}>
      <div className={`${styles['nav-items']} horizontal-scrollbar`}>
        <div className={styles['nav-item']}>
          <Link to='home' >
            <HomeIcon />
            Home
          </Link>
        </div>
        <div className={styles['nav-item']}>Neighbour</div>
        <div className={styles['nav-item']}>
          <Link to='landing-page/login'>
            <ForwardArrowIcon />
            Login
          </Link>
        </div>
        <div className={styles['nav-item']} onClick={() => logout()}>
          <Link to='logout'>
            <Icons name='BackArrowIcon'/>
            Logout
          </Link>
        </div>
      </div>
    </div>
  )
}

export default Navbar;