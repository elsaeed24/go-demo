package main // الحزمة الأساسية المطلوبة لتشغيل البرنامج

import "fmt" // استيراد مكتبة fmt الخاصة بالطباعة والإدخال

func main() { // الدالة الرئيسية التي يبدأ منها تنفيذ البرنامج

	// تعريف متغير نصي لتخزين الاسم
	var name string

	// تعريف متغير رقمي صحيح لتخزين العمر
	var age int

	// تعريف متغير عشري لتخزين الدرجة
	var grade float64

	// تعريف متغير منطقي (صح أو خطأ)
	var isPassed bool

	// تعريف Array تحتوي على 3 مواد
	var subjects [3]string

	// تعريف متغير لتخزين الأخطاء
	var err error

	// طباعة عنوان البرنامج
	fmt.Println("===== Student Profile =====")

	// طلب الاسم من المستخدم
	fmt.Print("Enter Name: ")

	// قراءة الاسم من المستخدم وتخزينه في name
	fmt.Scan(&name)

	// طلب العمر من المستخدم
	fmt.Print("Enter Age: ")

	// قراءة العمر من المستخدم
	// "_" يعني تجاهل أول قيمة راجعة
	// err تستقبل الخطأ إن وجد
	_, err = fmt.Scan(&age)

	// التحقق هل حدث خطأ أثناء الإدخال
	if err != nil {

		// طباعة رسالة الخطأ
		fmt.Println("Invalid age:", err)

		// إنهاء البرنامج
		return
	}

	// طلب الدرجة
	fmt.Print("Enter Grade: ")

	// قراءة الدرجة من المستخدم
	fmt.Scan(&grade)

	// سؤال هل الطالب ناجح
	fmt.Print("Passed? (true/false): ")

	// قراءة true أو false
	fmt.Scan(&isPassed)

	// طلب إدخال المواد
	fmt.Println("Enter 3 subjects:")

	// Loop يتكرر بعدد المواد داخل الـ Array
	for i := 0; i < len(subjects); i++ {

		// عرض رقم المادة
		fmt.Printf("Subject %d: ", i+1)

		// تخزين المادة داخل مكانها في Array
		fmt.Scan(&subjects[i])
	}

	// سطر جديد + عنوان النتائج
	fmt.Println("\n===== Student Data =====")

	// طباعة البيانات المدخلة
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Grade:", grade)
	fmt.Println("Passed:", isPassed)

	// عنوان المواد
	fmt.Println("Subjects:")

	// Loop يمر على كل عنصر داخل الـ Array
	for _, subject := range subjects {

		// طباعة اسم المادة
		fmt.Println("-", subject)
	}
}
