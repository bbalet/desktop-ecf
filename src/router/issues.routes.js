import { Layout, List, AddEdit } from '@/views/issues';

export default {
    path: '/issues',
    component: Layout,
    children: [
        { path: '', component: List },
        { path: 'add', component: AddEdit },
        { path: 'edit/:id', component: AddEdit }
    ]
};
