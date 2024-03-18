import { createMemoryHistory, createRouter } from 'vue-router'

import PedidosView from './views/PedidosView.vue'
// import ItemsView from './views/ItemsView.vue'

const routes = [
  { path: '/pedidos', component: PedidosView },
  // { path: '/items', component: ItemsView },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router
