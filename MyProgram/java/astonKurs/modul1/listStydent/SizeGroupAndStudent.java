import java.util.Arrays;

public final class SizeGroupAndStudent {
    private final String[] NAME_GROUP;
    private final String[] NAME_STUDENT;
    private final String[] SURNAME_STUDENT;
    private final String[] GENDER_STUDENT;
    private final int[] ADG_STUDENT;
    private final int[] QUANTITY_STUDENT_OR_GROUP;
    private final int QUANTITY_GROUP;


    public SizeGroupAndStudent(String[] nameGroup) {
        this.NAME_GROUP = Arrays.copyOf(nameGroup, nameGroup.length);
//        this.NAME_STUDENT = Arrays.copyOf(nameStudent, nameStudent.length);
//        this.SURNAME_STUDENT = Arrays.copyOf(surnameStudent, surnameStudent.length);
//        this.GENDER_STUDENT = Arrays.copyOf(genderStudent, genderStudent.length);
//        this.ADG_STUDENT = Arrays.copyOf(adgStudent, adgStudent.length);
//        this.QUANTITY_STUDENT_OR_GROUP = Arrays.copyOf(quantityStudentInGroup, quantityStudentInGroup.length);
//        this.QUANTITY_GROUP = quantityGroups;
    }
    public SizeGroupAndStudent(String[] nameStudent) {
        this.NAME_STUDENT = Arrays.copyOf(nameStudent, nameStudent.length);
    }
    public SizeGroupAndStudent(String[] surnameStudent) {
        this.SURNAME_STUDENT = Arrays.copyOf(surnameStudent, surnameStudent.length);
    }



    public String[] getNAME_GROUP() {
        return Arrays.copyOf(NAME_GROUP, NAME_GROUP.length);
    }
    public String[] getNAME_STUDENT() {
        return Arrays.copyOf(NAME_STUDENT, NAME_STUDENT.length);
    }
    public String[] getSURNAME_STUDENT() {
        return Arrays.copyOf(SURNAME_STUDENT, SURNAME_STUDENT.length);
    }
    public String[] getGENDER_STUDENT() {
        return Arrays.copyOf(GENDER_STUDENT, GENDER_STUDENT.length);
    }
    public int[] getADG_STUDENT() {
        return Arrays.copyOf(ADG_STUDENT, ADG_STUDENT.length);
    }
    public int[] getQUANTITY_STUDENT_OR_GROUP() {
        return Arrays.copyOf(QUANTITY_STUDENT_OR_GROUP, QUANTITY_STUDENT_OR_GROUP.length);
    }
    public int getQUANTITY_GROUP() {
        return QUANTITY_GROUP;
    }
}