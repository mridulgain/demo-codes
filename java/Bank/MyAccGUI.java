import javax.swing.JOptionPane;

public class MyAccGUI{
    public static void main(String[] args){
        boolean flag = true;
        MyAccount acc = null;
        while(flag){
            int op = JOptionPane.showConfirmDialog(
                            null, 
                            "Would you like to create an account?",
                            "Create account",
                            JOptionPane.YES_NO_OPTION,
                            JOptionPane.QUESTION_MESSAGE);
            // System.out.println(result);
            switch(op){
                case JOptionPane.YES_OPTION:
                    String name = JOptionPane.showInputDialog("What is your name?");
                    acc = new MyAccount(name);
                    String message = "Account creation successful";
                    message += "\nAccount Number: " + acc.accNo;
                    message += "\nAccount holder: " + acc.accHoldersName;
                    JOptionPane.showMessageDialog(null, message); 
                    break;
                case JOptionPane.NO_OPTION:
                    flag = false;
            }
        }
    }
}
