package myProgGui;

import javax.swing.*;

public class Window {

    public static void guiWin(String nameWin, String lableText) {
        JFrame frame = new JFrame(nameWin);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        JLabel label = new JLabel(lableText);
        frame.getContentPane().add(label);

        frame.pack();
        frame.setVisible(true);
    }
}
