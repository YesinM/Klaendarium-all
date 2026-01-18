import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    loaded: false,
    username: null,
    isAdmin: false,
  }),

  getters: {
    isLoggedIn: (state) => !!state.username,
  },

  actions: {
    async loadUser() {
      const res = await fetch('/api/getUserID', {
        credentials: 'include',
      })

      if (res.ok) {
        const data = await res.json()
        this.username = data.username
        this.isAdmin = data.isAdmin
      } else {
        this.$reset()
      }

      this.loaded = true
    },

    async logout() {
      await fetch('/api/logout', {
        method: 'POST',
        credentials: 'include',
      })
      this.$reset()
      this.loaded = true
    },
  },
})
