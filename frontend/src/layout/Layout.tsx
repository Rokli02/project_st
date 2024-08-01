import { FC } from 'react'
import { Route, Routes } from 'react-router-dom'
import Logout from './Logout'
import HomePage from '../pages/Home/Home'
import LoginPage from '../pages/LandingPage/Login/Login'
import SignUpPage from '../pages/LandingPage/SignUp/SignUp'
import Navbar from '../components/Navbar/Navbar'
import ProtectedRoutes from './ProtectedRoutes'

export const Layout: FC = () => {
  return (
    <div className='app-layout'>
      <Navbar />
      <Routes>
        <Route path='/' Component={ProtectedRoutes}>
          <Route index path='/home' Component={HomePage} />
          <Route path='/logout' Component={Logout}/>
        </Route>
        <Route path='/login' Component={LoginPage} />
        <Route path='/sign-up' Component={SignUpPage} />
      </Routes>
    </div>
  )
}

export default Layout