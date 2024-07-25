import { useCallback, useState } from "react"
import User from "../../models/common/User"

export const useHandleUser = () => {
  const [userData, setUserData] = useState<User | null>(null)

  const login = useCallback((login: string, password: string) => {
    // LOGIC

    setUserData({ login })
  }, [])

  const logout = useCallback(() => {
    setUserData(null)
  }, [])

  return {
    userData,
    login,
    logout,
  }
}