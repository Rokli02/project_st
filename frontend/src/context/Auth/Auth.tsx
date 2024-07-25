import { createContext, FC, ReactNode, useCallback, useContext } from "react";
import User from "../../models/common/User";
import { useHandleUser } from "./useHandleUser";

export interface IAuthContext {
  userData: User | null;
  hasUser: () => boolean;
  login: (username: string, password: string) => void;
  logout: () => void;
}

export interface IAuthProps {
  children: ReactNode;
}

const AuthContext = createContext<IAuthContext | null>(null)

export const AuthProvider: FC<IAuthProps> = ({ children }) => {
  const { userData, login, logout } = useHandleUser();

  const hasUser = useCallback(() => {
    return userData != null
  }, [])

  return (
    <>
    <AuthContext.Provider
      value={{
        userData,
        login,
        logout,
        hasUser,
      }}
    >
      {children}
    </AuthContext.Provider>
    </>
  )
}

export const useAuth = () => useContext(AuthContext) as IAuthContext

export default useAuth;