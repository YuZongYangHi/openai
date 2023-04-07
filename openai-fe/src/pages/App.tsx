import {useEffect, useState} from "react";
import {useLocation} from 'umi'
import {fetchMessageList, createMessage} from '@/services/openai/openai'
import { Comment } from '@ant-design/compatible';
import {Avatar, Input, Form, Button, Divider, Space, List} from 'antd';
import {PageContainer, ProCard} from '@ant-design/pro-components'
import moment from "moment";

const { TextArea } = Input;

let beforeDatetime = ""

const handleRenderDeadLines = (datetime: any) => {
  if (beforeDatetime.length === 0) {
    beforeDatetime = datetime
    return <Divider plain>{moment(datetime).format('YYYY-MM-DD HH:mm')}</Divider>
  }

  const elem = moment(datetime).diff(moment(beforeDatetime),'days') >= 1 &&  <Divider plain >{moment(datetime).format('YYYY-MM-DD HH:mm')}</Divider> || "";
  beforeDatetime = datetime;
  return elem
}

interface EditorProps {
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  onSubmit: () => void;
  submitting: boolean;
  value: string;
}

const Editor = ({ onChange, onSubmit, submitting, value }: EditorProps) => (
  <>
    <Form.Item>
      <TextArea rows={4} onChange={onChange} value={value} />
    </Form.Item>
    <Form.Item>
      <Button htmlType="submit" loading={submitting} onClick={onSubmit} type="primary" style={{float: "right"}}>
        Send Message
      </Button>
    </Form.Item>
  </>
);

export default () => {
  const [channelId, setChannelId] = useState(0);
  const [messageList, setMessageList] = useState([]);
  const [submitting, setSubmitting] = useState(false);
  const [value, setValue] = useState('');
  const [loading, setLoading] = useState(true);
  const location = useLocation();

  const handleChange = (e) => {
    beforeDatetime = ""
    setValue(e.target.value)
  }
  const handleSubmit = async () => {
    if (value.length === 0) {
      return
    }
    const v = value
    setValue("");
    let merge = [...messageList, {
      id: 100000,
      channelId: channelId,
      content: v,
      dialogType: 1,
      createdAt: moment().format('YYYY-MM-DD HH:mm')
    }]
    setMessageList(merge)
    setSubmitting(true);

    const result = await createMessage(channelId, v)
    if (result.code === 200) {
      merge = [...merge, result.data]
      setMessageList(merge)
      setSubmitting(false)
    }
  }

  useEffect(()=>{

    (async function init() {
      setLoading(true)
      const cid = parseInt(location.pathname.split('/')[2])
      setChannelId(cid)
      const result = await fetchMessageList(cid)
      if (result.code === 200) {
        setMessageList(result.data);
      }
      setLoading(false)
    })()
    }, [location.pathname])

  return (
    <PageContainer
      ghost
      title={false}
      loading={loading}
    >

      <ProCard title={false}  bordered>

        <List
          className="comment-list"
          itemLayout="horizontal"
          dataSource={messageList}
          renderItem={item => (
            <>
              {
                handleRenderDeadLines(item.createdAt)
              }
              {item.dialogType === 1 ?
                (
                  <div className="space-align-block" key={item.id} style={{marginBottom: 32, height: 40}}>
                    <Space align="center" style={{float: "right", clear: 'both'}}>
              <span style={{
                padding: 8,
                color: "#222226",
                background: "#cad9ff",
                display: "inline-block"
              }}>{item.content}</span>
                      <Avatar src={"/user.png"}/>
                    </Space>
                  </div>
                ) :
                (

                  <div>
                    <Space>
                    <Comment
                      author={"ChatGPT"}
                      avatar={"/logo.svg"}
                      content={<span style={{
                        padding: 8,
                        display: "inline-block",
                        color: "#8b4513",
                        maxWidth: "800px",
                        background: "#f5f6f7",
                        boxShadow: "0 1px 2px 0 rgb(0 0 0 / 10%)",
                        whiteSpace: 'pre-line',
                      }}>{item.content}</span>}
                      datetime={moment(item.createdAt).format('YYYY-MM-DD HH:mm')}
                    />
                    </Space>
                  </div>
                )
              }
            </>
          )}
        />

        <Comment
          avatar={<Avatar src="/user.png" alt="user" />}
          content={
            <Editor
              onChange={handleChange}
              onSubmit={handleSubmit}
              submitting={submitting}
              value={value}
            />}
          />
      </ProCard>
    </PageContainer>
  )
}
