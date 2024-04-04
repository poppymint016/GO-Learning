## Overview
โปรเจ็คนี้จัดทำเพื่อเป็นส่วนหนึ่งของการเรียนรู้ Golang Programing ซึ่งเป็นโปรเจคเกี่ยวกับการเพิ่มประสบการณ์การทำงาน ประกอบไปด้วย Method Post/Put/Get/Delete โปรเจคนี้เป็นเพียงพื้นฐานการเขียน Rest-Api เท่านั้น โดยที่โปรเจคจะเขียนในโครงสร้าง Architecture Business Logic Layer หรือ Clean Architecture พร้อมกับการเขียน Unit test 

### Golang 
- Framework --> fiber
- Test Package--> testify / Uber-Mock
- Database --> MongoDB

### API Structure
1. Create Experience
- Method POST /api/experiences
- สร้างข้อมูลประสบการณ์
2. Update Experince
- Method PUT /api/:experienceId
- แก้ไขข้อมูลประสบการณ์
3. FindById Experience
- Method GET /api/:experienceId
- ดูข้อมูลประสบการณ์แบบเฉพาะเจาะจง
4. FindById Experience
- Method GET /api
- ดูข้อมูลประสบการณ์ทั้งหมด
5. Delete Experience
- Method DELETE /api
- ลบข้อมูลประสบการณ์
