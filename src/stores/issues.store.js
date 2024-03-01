import { defineStore } from 'pinia'
import { fetchWrapper } from '@/helpers'
import { globals } from '../main'

const baseApiUrl = `${globals.baseApiUrl}/issues`

export const useUsersStore = defineStore({
    id: 'issues',
    state: () => ({
        issues: {},
        issue: {}
    }),
    actions: {
        async getAll() {
            this.issues = { loading: true };
            try {
                this.issues = await fetchWrapper.get(baseApiUrl);    
            } catch (error) {
                this.issues = { error };
            }
        },
        async getById(id) {
            this.issue = { loading: true };
            try {
                this.issue = await fetchWrapper.get(`${baseApiUrl}/${id}`);
            } catch (error) {
                this.issue = { error };
            }
        },
        async update(id, params) {
            await fetchWrapper.put(`${baseApiUrl}/${id}`, params);

            
        },
    }
});
