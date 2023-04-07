import { request } from 'umi';

export async function fetchChannelList() {
  return request('/api/v1/openai/channel', {
    method: 'GET',
  })
}

export async function fetchChannelObj(channelId: number) {
  return request(`/api/v1/openai/channel/${channelId}`, {
    method: 'GET',
  })
}

export async function updateChannel(channelId: number, title: string) {
  const data = {
    title: title
  }
  return request(`/api/v1/openai/channel/${channelId}`, {
    method: 'POST',
    data: data
  })
}

export async function deleteChannel(channelId: number) {
  return request(`/api/v1/openai/channel/${channelId}`, {
    method: 'DELETE',
  })
}

export async function createChannel(title: string) {
  const data = {
    title: title
  }
  return request(`/api/v1/openai/channel`, {
    method: 'POST',
    data: data
  })
}


export async function fetchMessageList(channelId: number) {
  return request(`/api/v1/openai/message/${channelId}`, {
    method: 'GET',
  })
}

export async function createMessage(channelId: number, content: string) {
  const data = {
    content: content
  }
  return request(`/api/v1/openai/message/${channelId}`, {
    method: 'POST',
    data: data
  })
}

