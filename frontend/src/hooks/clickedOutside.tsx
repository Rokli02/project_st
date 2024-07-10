import { RefObject, useEffect } from "react";

type FICO = (
  elementRef: RefObject<HTMLElement | null>,
  shouldTrigger: boolean,
  onClickedOutside: () => void,
) => void

export const useClickedOutside: FICO = (elementRef, shouldTrigger = true, onClickedOutside) => {
  useEffect(
    () => {
      if (!shouldTrigger) return

      function clickedOutside(ev: MouseEvent) {
        if (!elementRef.current || !elementRef.current.contains(ev.target as any)) {
          onClickedOutside()
        }
      }

      document.addEventListener('mousedown', clickedOutside)

      return () => {
        document.removeEventListener('mousedown', clickedOutside)
      }
    },
    [elementRef, shouldTrigger]
  )
}

export default useClickedOutside;