
# Функциональность микросервисов

1. practice-task-service - микросервис, отвечающий за загрузку, получение и удаления **заданий по практическим работам** от преподавателей.

    Его методы:
    1. Получение всех практических
        GET /api/practice-tasks/

        {
            json: data
        }
    
    2. Получение одной практической
  
        GET /api/practice-tasks/{id}

        {
           json: data
        }
       
    3. Поиск практических работ
  
       GET /api/practice-tasks/search?title="название"&item="предмет"

       {
            json: data
       }

    4. Загрузка практической работы
  
       POST /api/practice-tasks

       {
            json: data
       }
    
    5. Удаление практической работы
  
       DELETE /api/practice-tasks/{id}

       {
            json: data
       }
