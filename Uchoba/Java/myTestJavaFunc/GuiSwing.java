import javax.swing.*;
import java.awt.*;

public class GuiSwing {

    public static void main(String[] args) {
        JFrame frame = new JFrame("Пример Swing");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setSize(400, 300);

        // 2. Создаем панель с менеджером компоновки
        JPanel panel = new JPanel(new BorderLayout(10, 10));
        panel.setBorder(BorderFactory.createEmptyBorder(10, 10, 10, 10));

        // 3. Создаем компоненты
        JLabel label = new JLabel("Введите имя:", SwingConstants.CENTER);
        label.setFont(new Font("Arial", Font.BOLD, 14));

        JTextField textField = new JTextField(20);

        JButton button = new JButton("OK");
        button.setPreferredSize(new Dimension(100, 30));
        button.addActionListener(e -> {
            JOptionPane.showMessageDialog(frame,
                    "Привет, " + textField.getText() + "!");
        });

        // 4. Панель для кнопок с FlowLayout
        JPanel buttonPanel = new JPanel(new FlowLayout(FlowLayout.CENTER));
        buttonPanel.add(button);

        // 5. Добавляем компоненты в основную панель
        panel.add(label, BorderLayout.NORTH);
        panel.add(textField, BorderLayout.CENTER);
        panel.add(buttonPanel, BorderLayout.SOUTH);

        // 6. Добавляем панель в окно и показываем
        frame.add(panel);
        frame.setVisible(true);
    }
}
