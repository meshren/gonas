import Layout from '@/layout'

export default {
  path: '/file',
  component: Layout,
  meta: { title: 'file.management', icon: 'tree-table' },
  children: [
    {
      path: 'index',
      name: 'file.index', // tab栏显示的名称
      component: () => import('@/views/file/index'),
      meta: { title: 'file.index', affix: false }
    },
    {
      path: 'upload',
      name: 'file.upload',
      component: () => import('@/views/file/upload'),
      meta: { title: 'file.upload', affix: false } // 左侧菜单栏显示的名称
    }
  ]
}
