import { defineStore } from 'pinia'
import { fetchWrapper } from '@/helpers'
import { baseApiUrl } from '../main'

export const useTheatersStore = defineStore({
    id: 'theaters',
    state: () => {
        return {
            theaters: [],
        }
    },
    actions: {
        async getAll() {
            try {
                this.theaters = await fetchWrapper.get(`${baseApiUrl}/api/theaters/`);    
            } catch (error) {
                this.theaters = { error };
            }
        },
    }
});
