class Student:
    def readFile(self):
        with open("../students.txt", 'r', encoding='utf-8') as f:
            students = f.read().splitlines()
        return students

    def process_grades(self, records: list[str]) -> dict:
        valid_records = []
        skipped = 0

        for record in records:
            if ": " not in record:
                skipped += 1
                continue

            parts = record.split(": ")
            if len(parts) != 2:
                skipped += 1
                continue

            name, score_str = parts
            name = name.strip()

            if name == "" or not score_str.isdigit():
                skipped += 1
                continue

            score = int(score_str)
            valid_records.append((name, score))

        valid_count = len(valid_records)

        if valid_count > 0:
            average = sum(score for _, score in valid_records) / valid_count
            average = round(average, 1)
        else:
            average = 0.0

        passed = [name for name, score in valid_records if score >= 60]
        passed.sort()
        return {
            "valid_count": valid_count,
            "average": average,
            "passed": passed,
            "skipped": skipped
        }

    def run(self):
        records = self.readFile()

        if not records:
            print("Нет данных для обработки")
            return None
        result = self.process_grades(records)
        print("Результат обработки:")
        print(f"Всего записей: {len(records)}")
        print(f"Валидных записей: {result['valid_count']}")
        print(f"Пропущено (битых): {result['skipped']}")
        print(f"Средняя оценка: {result['average']}")
        print(f"Студенты, сдавшие экзамен (>= 60): {result['passed']}")
        print(f"Количество сдавших: {len(result['passed'])}")

        return result

print(Student().run())