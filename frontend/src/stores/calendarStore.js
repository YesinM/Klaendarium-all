import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useCalendarFiltersStore = defineStore('calendarFilters', () => {
  const searchQuery = ref('')
  const chosedYears = ref([])

  function setSearchQuery(query) {
    searchQuery.value = query
  }

  function setChosedYears(years) {
    chosedYears.value = years
  }


  return {
    searchQuery,
    chosedYears,
    setSearchQuery,
    setChosedYears
  }
})
