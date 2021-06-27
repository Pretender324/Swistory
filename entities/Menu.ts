export type Group = {
  id: number
  menus: Menu[]
  name: string
}

export type Menu = {
  id: number
  level: string
  aim: string
  description: string
  createdAt: string
  group: Partial<Group>
  practices: Practice
  updatedAt: string
}

export type Practice = {
  category: String
  createdAt: string
  description: string
  distance: string
  id: number
  menu: Menu
  onTime: string
  set: number
  system: string
  times: number
  updatedAt: string
}
