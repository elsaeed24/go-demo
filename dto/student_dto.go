package dto

// CreateStudentInput Create DTO
type CreateStudentInput struct {
	Name      string `json:"name" binding:"required,min=3"` // لازم + أقل حاجة 3 حروف
	Age       int    `json:"age" binding:"required,min=1"`  // لازم + أكبر من 0
	TeacherID uint   `json:"teacher_id" binding:"required"`
}

// UpdateStudentInput Update DTO
type UpdateStudentInput struct {
	Name      string `json:"name" binding:"omitempty,min=3"` // optional
	Age       int    `json:"age" binding:"omitempty,min=1"`
	TeacherID uint   `json:"teacher_id"`
}
