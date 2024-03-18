import { createMemoryHistory, createRouter } from 'vue-router'
import { Chart } from "frappe-charts/dist/frappe-charts.min.esm" 
import PedidosView from './views/PedidosView.vue'
import ComprasView from './views/ComprasView.vue'
import InventariosView from './views/InventariosView.vue'

const routes = [
  {
    path: '/',
    children: [
      {
        path: '',
        component: PedidosView
      },
      {
        path: 'pedidos',
        component: PedidosView
      },
      {
        path: 'compras',
        component: ComprasView
      },
      {
        path: 'inventarios',
        component: InventariosView
      }
    ]
  }
]


const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router
