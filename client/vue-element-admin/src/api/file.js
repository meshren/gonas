import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/files',
    method: 'get',
    params: query
  })
}

export function fetchFile(id) {
  return request({
    url: '/files/' + id,
    method: 'get',
    params: { id }
  })
}

export function uploadFile(data) {
  return request({
    url: '/files',
    method: 'post',
    data
  })
}

export function createFolder(data) {
  return request({
    url: '/folders',
    method: 'post',
    data
  })
}
