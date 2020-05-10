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

export function updateArticle(data) {
  return request({
    url: '/vue-element-admin/article/update',
    method: 'post',
    data
  })
}
