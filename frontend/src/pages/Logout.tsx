
import { useLayoutEffect, FC } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context'

export const Logout: FC = () => {
  const { logout } = useAuth()
  const navigate = useNavigate()

  useLayoutEffect(() => {
    logout()
    navigate('/', { replace: true })
  }, [])

  return (
    <div>Logout</div>
  )
}
