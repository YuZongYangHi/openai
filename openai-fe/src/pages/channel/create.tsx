import {Row, Col, message, Space} from 'antd';
import {
  PageContainer,
  ProForm,
  ProFormText,
} from '@ant-design/pro-components';
import {createChannel} from "@/services/openai/openai";
import {history, useModel} from 'umi'
import {result} from "lodash";

const handleCreateChannel = async (values) =>{
  return await createChannel(values.title)
}

const formItemLayout = {
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },
}

export default () =>{
  return (
    <PageContainer title={false} ghost={true}>
    <ProForm<{
      name: string;
      company?: string;
      useMode?: string;
    }>
      {...formItemLayout}
      layout={"horizontal"}
      submitter={{
        render: (props, doms) => {
         return <Row>
            <Col span={14} offset={4}>
              <Space>{doms}</Space>
            </Col>
          </Row>
        },
      }}
      onFinish={async (values) => {
        const result = await handleCreateChannel(values);
        message.success('add success');
        window.location.href = `/channel/${result.data.id}`
      }}
    >
      <ProFormText
        width="md"
        name="title"
        label="title"
        placeholder="input title"
      />
    </ProForm>
    </PageContainer>
  )
}
