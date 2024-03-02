import { defineStore } from 'pinia'
import { fetchWrapper } from '@/helpers'
import { baseApiUrl } from '../main'

export const useIssuesStore = defineStore({
    id: 'issues',
    state: () => ({
        issues: {},
        issue: {}
    }),
    actions: {
        async getAll() {
            this.issues = { loading: true };
            try {
                this.issues = await fetchWrapper.get(`${baseApiUrl}/api/issues/`);    
            } catch (error) {
                this.issues = { error };
            }
        },
        async getById(id) {
            this.issue = { loading: true };
            try {
                this.issue = await fetchWrapper.get(`${baseApiUrl}/api/issues/${id}`);
            } catch (error) {
                this.issue = { error };
            }
        },
        async update(id, params) {
            await fetchWrapper.put(`${baseApiUrl}/api/issues/${id}`, params);

            
        },
    }
});
