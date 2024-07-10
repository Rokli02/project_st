export interface IconProps {
  className?: string;
}

export interface GeneralIconProps extends IconProps {
  name: string
}

export type IconComponent = (props: IconProps) => JSX.Element

export interface IconObject {
  [k: string]: IconComponent
}