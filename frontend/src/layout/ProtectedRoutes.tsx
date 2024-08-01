
import { FC } from 'react'
import { useAuth } from '../context';
import { Redirect } from './Redirect';
import { Outlet } from 'react-router-dom';

interface IProtectedRoutes {}

export const ProtectedRoutes: FC<IProtectedRoutes> = () => {
  const { hasUser } = useAuth()

  if (!hasUser) {
    return <Redirect to='/login'/>
  }

  return (
    <div style={{ backgroundColor: 'pink' }} >
      <Outlet />
    </div>
  )
}

export default ProtectedRoutes;
