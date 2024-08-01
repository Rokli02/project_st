
import { FC, useLayoutEffect } from 'react'
import { useAuth } from '../context'

export const Logout: FC = () => {
  const { logout } = useAuth()
  
  useLayoutEffect(() => {
    logout()
  })

  return <div>Logout</div>
}

export default Logout;
