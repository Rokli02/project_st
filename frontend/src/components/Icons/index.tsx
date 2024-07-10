export * from './heroicons_1';
import * as heroicons_1 from './heroicons_1';
import { GeneralIconProps, IconObject } from './types';

const _Icons = {
  ...heroicons_1,
}
const Icons = ({name, ...props}: GeneralIconProps) => (_Icons as IconObject)[name](props);

export default Icons;