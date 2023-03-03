import { createRouter, createWebHashHistory } from 'vue-router';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import Index from '../views/Index.vue'
import Home from '../views/Home.vue'
import Error from '../views/Error.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Pass from '../views/Pass.vue'
import Dashboard from '../views/Dashboard.vue'
import Student from '../views/Student.vue'
import Project from '../views/Project.vue'
import Output from '../views/Output.vue'
import Config from '../views/Config.vue'
import Subscribe from '../views/Subscribe.vue'
import Result from '../views/Result.vue'

const routes = [
  {
    path: '/',
    component: Index,
    redirect: '/login',
    children: [
      {
        path: '/login',
        component: Login,
      },
      {
        path: '/register',
        component: Register,
      },
      {
        path: '/pass',
        component: Pass,
      },
    ],
  },
  {
    path: '/home',
    component: Home,
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
        name: 'dashboard',
        component: Dashboard,
      },
      {
        path: '/student',
        name: 'student',
        component: Student,
      },
      {
        path: '/project',
        name: 'project',
        component: Project,
      },
      {
        path: '/output',
        name: 'output',
        component: Output,
      },
      {
        path: '/config',
        name: 'config',
        component: Config,
      },
      {
        path: '/subscribe',
        name: 'subscribe',
        component: Subscribe,
      },
      {
        path: '/result',
        name: 'result',
        component: Result,
      }
    ],
  },
  {
    path: '/error',
    name: 'error',
    component: Error,
  },
]

const router = createRouter({
  history: createWebHashHistory(), routes
})

NProgress.configure({ easing: 'ease', speed: 500, showSpinner: false });

router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})

router.afterEach(() => {
  NProgress.done()
})

export default router;