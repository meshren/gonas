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

export function fetchDirectory(parent_id) {
  return request({
    url: '/directories',
    method: 'get',
    params: { parent_id }
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
    url: '/directories',
    method: 'post',
    data
  })
}
