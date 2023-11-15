
# Функциональность микросервисов

* practice-task-service - микросервис, отвечающий за загрузку, получение и удаления **заданий по практическим работам** от преподавателей.

    Его методы:
    1. Получение всех практических

           GET /api/practice-tasks/

            {
                json: data
            }
    
    3. Получение одной практической
  
            GET /api/practice-tasks/{id}

            {
               json: data
            }
       
    4. Поиск практических работ
  
           GET /api/practice-tasks/search?title="название"&item="предмет"

           {
               json: data
           }

    5. Загрузка практической работы
  
           POST /api/practice-tasks

           {
               json: data
           }
    
    6. Удаление практической работы
  
           DELETE /api/practice-tasks/{id}

           {
               json: data
           }
