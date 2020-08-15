export interface MenuItem {
  icon: string;
  text: string;
  route?: string[];
  children?: MenuItem[];
}