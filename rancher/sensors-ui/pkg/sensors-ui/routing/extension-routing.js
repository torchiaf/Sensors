// Don't forget to create a VueJS page called index.vue in the /pages folder!!!
import Dashboard from '../pages/index.vue';

const BLANK_CLUSTER = '_';
const YOUR_PRODUCT_NAME = 'sensors';

const routes = [
  {
    route: {
      name: `${YOUR_PRODUCT_NAME}-c-cluster`,
      path: `/${YOUR_PRODUCT_NAME}/c/:cluster`,
      component: Dashboard,
      meta: {
        product: YOUR_PRODUCT_NAME,
        cluster: BLANK_CLUSTER,
        pkg: YOUR_PRODUCT_NAME,
      },
    },
    parent: 'plain'
  },
];

export default routes;