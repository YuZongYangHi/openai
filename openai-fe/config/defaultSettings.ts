import { ProLayoutProps } from '@ant-design/pro-components';

/**
 * @name
 */
const Settings: ProLayoutProps & {
  pwa?: boolean;
  logo?: string;
} = {
  "navTheme": "light",
  "colorPrimary": "#13C2C2",
  "layout": "mix",
  "contentWidth": "Fluid",
  "fixedHeader": true,
  "fixSiderbar": true,
  "pwa": true,
  "logo": "/logo.svg",
  "title": "ChatGPT",
  "token": {},
  "splitMenus": false,
  "siderMenuType": "sub",
  "menu": {
    "locale": false
  }
}

export default Settings;
