import Footer from '@/components/Footer';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from '@umijs/max';
import defaultSettings from '../config/defaultSettings';
import { errorConfig } from './requestErrorConfig';
import React, {useEffect, useState} from 'react';
import { AvatarDropdown, AvatarName } from './components/RightContent/AvatarDropdown';
import {fetchChannelList} from "@/services/openai/openai";
import App from '@/pages/App'
import {MessageOutlined} from '@ant-design/icons'
import Welcome from "@/pages/Welcome";
import {Button, MenuProps, Dropdown, Popconfirm, message} from 'antd'
import {Link, useModel, history} from 'umi'
import {deleteChannel} from "@/services/openai/openai";
import CreateChannel from '@/pages/channel/create'
import routes from "../config/routes";

let extraRoutes;
let routeIndex = 0

export function patchRoutes({ routes, routeComponents }) {
}

export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  loading?: boolean;
}> {
  return {
    settings: defaultSettings as Partial<LayoutSettings>,
  };
}

export const patchClientRoutes = ({ routes }) => {
  routes.forEach((item, index)=>{
    if (item.id === "ant-design-pro-layout") {
      routes[index].children = []
      routeIndex = index
      extraRoutes.forEach(extraRoute => {
        routes[index].children.push({
          element: <App/>,
          path: `/channel/${extraRoute.id}`,
          name: extraRoute.title,
          icon: <MessageOutlined />,
          id: extraRoute.id,
          parentId: "ant-design-pro-layout",
        })
      })

      routes[index].children.push({
        element: <CreateChannel/>,
        path: "/channel/dispatch/create",
        hideInMenu: true,
        key: "create_channel",
        id: "create_channel",
        parentId: "ant-design-pro-layout"
      })

      routes[index].children.push({
        element: <Welcome/>,
        path: "/",
        hideInMenu: true,
        key: "default",
        id: "default",
        parentId: "ant-design-pro-layout"
      })
    }
  })
};

export function render(oldRender) {
  fetchChannelList().then(res=>{
    extraRoutes = res.data
    oldRender();
  })
}

export const layout: RunTimeLayoutConfig = ({ initialState, setInitialState }) => {
  const [channelId, setChannelId] = useState(0);
  const handleDeleteChannel = async (id: number) => {
    const result = await deleteChannel(id)
    if (result.code === 200) {
      message.success("delete success!")
      window.location.href = "/"
    }

  }
  const items: MenuProps['items'] = [
    {
      label: (
        <Popconfirm
          placement="top"
          title="Delete the channel"
          description="Are you sure to delete this channel?"
          okText="Yes"
          cancelText="No"
          onClick={e => e.stopPropagation()}
          onConfirm={async()=>{await handleDeleteChannel(channelId)}}
          getPopupContainer={trigger => trigger.parentNode}
        >
          <div>Delete</div>
        </Popconfirm>
      ),
      key: 'delete',
      danger: true,
    },

  ];

  return {
    avatarProps: {
      src: "igolang.cn",
      title: <AvatarName />,
      render: (_, avatarChildren) => {
        return <AvatarDropdown>{avatarChildren}</AvatarDropdown>;
      },
    },
    waterMarkProps: {
      content: "igolang.cn",
    },
    footerRender: () => <Footer />,
    onPageChange: () => {
    },
    menuItemRender: (item, dom) => {
      return <Dropdown
                onOpenChange={()=>{setChannelId(item.id)}}
                overlayStyle={{width: "180px"}}
                menu={{items}}
                placement="bottom"
                trigger={['contextMenu']}>
            <Link to={item.itemPath}>{dom}</Link>
      </Dropdown>
    },
    menuFooterRender: (props) => {
      return !props.collapsed && (
        <Button type="dashed" block onClick={()=>{
          history.push("/channel/dispatch/create")
        }}>Add New Channel</Button>
      )
    },
    layoutBgImgList: [],
    links: [],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
        </>
      );
    },
    ...initialState?.settings,
  };
};

export const request = {
  ...errorConfig,
};
