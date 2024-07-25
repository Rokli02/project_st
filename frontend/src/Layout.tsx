import { FC, useLayoutEffect } from 'react'
import { Route, Routes, useNavigate } from 'react-router-dom'
import { Logout } from './pages/Logout'
import HomePage from './pages/Home/Home'
import LoginPage from './pages/LandingPage/Login/Login'
import SignUpPage from './pages/LandingPage/SignUp/SignUp'
import Navbar from './components/Navbar/Navbar'
import { useAuth } from './context'

export const Layout: FC = () => {
  const { hasUser } = useAuth()

  useLayoutEffect(() => {
    if (hasUser())
      console.log('van bejelentkezett felhasználó')
    else
      console.log('nincs bejelentkezett felhasználó')
  }, [])

  return (
    <div className='app-layout'>
      <Navbar />
      <Routes>
        <Route index Component={BasePageSelector} />
        <Route path='landing-page' >
          <Route path='login' Component={LoginPage} />
          <Route path='sign-up' Component={SignUpPage} />
        </Route>
        <Route path='home' Component={HomePage} />
        <Route path='logout' Component={Logout}/>
      </Routes>
    </div>
  )
}

const BasePageSelector: FC = () => {
  const { hasUser } = useAuth()
  const navigate = useNavigate()

  if (hasUser()) {
    navigate('/home', { replace: true })
  } else {
    navigate('/landing-page/login')
  }

  console.log('BASE_PAGE')

  return <></>
}

export default Layout