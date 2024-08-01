import { FC, useLayoutEffect } from 'react'
import { useNavigate } from 'react-router-dom'

interface IRedirect {
  to: string;
  replace?: boolean;
}

export const Redirect: FC<IRedirect> = ({ to, replace = true }) => {
  const navigate = useNavigate()

  useLayoutEffect(() => {
    navigate(to, { replace })
  })

  return (
    <div></div>
  )
}

export default Redirect;
