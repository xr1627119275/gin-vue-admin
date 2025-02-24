export interface Task {
  id: number
  title: string
  description: string
  status: 'pending' | 'in_progress' | 'completed'
  assignee: string
  created_at: Date
  started_at?: Date
  completed_at?: Date
  priority: 'high' | 'medium' | 'low'
  ai_summary?: string
}

export interface Annotation {
  id: number
  user: string
  content: string
  timestamp: Date
}

export interface Alert {
  id: number
  severity: 'high' | 'medium' | 'low'
  device: string
  description: string
  timestamp: Date
  annotations?: Annotation[]
}